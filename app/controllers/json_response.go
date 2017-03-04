/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 02/03/2017    
* Time      : 22:56
* Developer : ibrahimcobani
*
*******/
package controllers

type JsonResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}
