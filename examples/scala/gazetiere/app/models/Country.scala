package models

import play.api.db._
import play.api.Play.current

import anorm._
import anorm.SqlParser._

case class Country(id: Int, name: String, path: String, source: String, groupType: String, _type: String, lat: Double, lon: Double,info: String)

object Country {

  def list: List[Country] = DB.withConnection { implicit session =>

    val result:List[Int~String~String~String~String~String~String~String~String] = {
      SQL("select * from address where type='country'")
        .as(
          get[Int]("id")
            ~get[String]("name")
            ~get[String]("path")
            ~get[String]("src")
            ~get[String]("group_type")
            ~get[String]("type")
            ~get[String]("lon")
            ~get[String]("lat")
            ~get[String]("info")*
        )
    }

    result.map{
      case id ~ name ~ path ~ source ~ groupType ~ _type ~ lon ~ lat ~ info =>
        Country(id, name, path, source, groupType, _type, lon.toDouble, lat.toDouble, info)
    }

  }

}
