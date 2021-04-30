/**
 * the json of the message it should to have this structure
 * ```
 * interface {
	Author  string `json:"author"`
	Message string `json:"message"`}
 */
interface Message{
    author:string;
    message:string;
}
///ws/{name:[\w\W]}
const conn=new WebSocket(`ws://${window.location.host}/ws/${window.location.pathname.split("/")[1]}`)
conn.addEventListener("message",(event)=>{
    
    console.log(JSON.parse(event.data))
})
