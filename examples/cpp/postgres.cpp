#include <evhttp.h>
#include <iostream>
#include "jsoncons/json.hpp"
#include <string>
#include <sstream>

#include "db.h"

using namespace std;
using jsoncons::json;

std::string getCountriesData(DatabasePQ* db)
{
	cout << "Fetching data from PQ" << endl;

	auto queryResult = db->query("SELECT id,name,path,src,group_type,type,lat,lon,info FROM address WHERE type='country' ORDER BY name");

	const char* fields[] = 
	{
		"id",
		"name",
		"path",
		"source",
		"group_type",
		"type",
		"lat",
		"lon",
		"info"
	};

	auto columnsCount = queryResult.columnsCount();
	auto rowsCount = queryResult.rowsCount();

	json countries(json::an_array);

	for (int i = 0; i < rowsCount; ++i)
	{
		json country;
		for (int j = 0; j < columnsCount; ++j)
		{
			auto value = queryResult.get(i, j);
			country[fields[j]] = std::string(value);
		}
		countries.add(country);
	}
	
	stringstream stream;
	stream << pretty_print(countries);

	return stream.str();
}

void onRequest(evhttp_request *req, void * arg)
{
	DatabasePQ* db = (DatabasePQ*)arg;

	evhttp_add_header(
		evhttp_request_get_output_headers(req),
		"Content-Type",
		"application/json");

	//get json data
	const auto& data = getCountriesData(db);

	auto dataBuf = evhttp_request_get_output_buffer(req);

	evbuffer_add_printf(
		dataBuf,
		data.c_str());

	evhttp_send_reply(req, HTTP_OK, "", dataBuf);
}

int main(int argc, char* argv[])
{
    DatabasePQ db;
    db.connect("host=osm-db-dev.srv.pv.km dbname=address user=reader password=reader");

    if (!event_init())
    {
    	cerr << "Unable to init libevent" << endl;
    	return -1;
    }

    auto server = evhttp_start("127.0.0.1", 8080);

    if (!server)
    {
    	cerr << "Unable to start local server" << endl;
    	return -1;
    }

    //set request callback
    evhttp_set_gencb(server, onRequest, (void*)&db);
    //run loop
    cout << "Running server..." << endl;
    event_dispatch();
    evhttp_free(server);
    return 0;
}