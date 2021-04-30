/**
 * the json of the message it should to have this structure
 * ```
 * interface {
	Author  string `json:"author"`
	Message string `json:"message"`}
 */
interface Message {
  author: string;
  message: string;
}
///ws/{name:[\w\W]}
console.log(window.location.pathname.split("/")[2]);
const ws = new WebSocket(
  `ws://${window.location.host}/ws/${window.location.pathname.split("/")[2]}`
);

const sendMsg=()=>{
    ws.send(JSON.stringify({
   author: document.getElementById("author" ).value,
   message:document.getElementById("message").value})
    

}
ws.onmessage=(event) => {
  console.log(JSON.parse(event.data));
}
