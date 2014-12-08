package controllers

import play.api.libs.json.Json
import play.api.mvc._
import models.Country

object Application extends Controller {

  def index = Action {

    val countryList = Country.list.map{
      c => Json.obj(
        "id" -> c.id,
        "name" -> c.name,
        "path" -> c.path,
        "position" -> Json.obj(
          "lat" -> c.lat,
          "lon" -> c.lon
        ),
        "src" -> c.source,
        "type" -> c._type,
        "group_type" -> c.groupType,
        "info" -> Json.parse(c.info)
      )
    }

    Ok(Json.toJson(countryList))

  }

}