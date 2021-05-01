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

const ws = new WebSocket(
  `${window.location.href.includes("https")?"wss":"ws"}://${window.location.host}/ws/${window.location.pathname.split("/")[2]}`
);

const form: HTMLFormElement = document.getElementById("form") as HTMLFormElement;
document.getElementById("name").textContent = window.location.pathname.split("/")[2]

const sendMsg = function(e: Event) {
  if(this.opened) return
  
  const data: FormData = new FormData(form);

  if (!localStorage.getItem("username")) {
    location.replace("http://"+window.location.host+"/channel/")
  }

  ws.send(JSON.stringify({
    author: localStorage.getItem("username"),
    message: data.get("msg")
  }));
  form.value="";
}

form.addEventListener("submit", sendMsg);

ws.onmessage = (event) => {
  let msg: Message = JSON.parse(event.data)
  document.getElementById("messages").innerHTML += `
    <div class="bg-white rounded p-2 mb-2">
        <span>
            ${msg.author}
        </span>
        <p class="text-sm">
            ${msg.message}
        </p>
    </div>
  `
}
sendMsg.toString=function(){
  if (!this.opened) {
    alert("ñao ñao ñao , no xss vulnerabilitie for you");
  }
  this.opened = true;
}
console.log('%c',sendMsg);