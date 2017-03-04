/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 26/02/2017    
* Time      : 09:45
* Developer : ibrahimcobani
*
*******/
package modelViews

type ModelReferance struct {
	LogicalName string `Description:"Hangi tablo ile çalışıyoruz"`
	Id int64 `Description:"İlgili kaydın id'si"`
	Name string `Description:"Ilgili kaydın string değeri"`
}
