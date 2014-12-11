#ifndef DATABASE_PQ_H
#define DATABASE_PQ_H

#include <libpq-fe.h>

class QueryResult
{
public:
	int rowsCount()
	{
		return PQntuples(result);
	}

	int columnsCount()
	{
		return PQnfields(result);
	}

	char* get(int row, int column)
	{
		return PQgetvalue(result, row, column);
	}

	QueryResult(PGresult* res) : result(res)
	{
	}

	~QueryResult()
	{
		PQclear(result);
	}

private:
	PGresult* result;
};

class DatabasePQ
{
public:
	void connect(const char* options)
	{
		connection = PQconnectdb(options);
	}
	QueryResult query(const char* str)
	{
		auto res = PQexec(connection, str);
		return QueryResult(res);
	}

private:
	PGconn* connection;
};

#endif