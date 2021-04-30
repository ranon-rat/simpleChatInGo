const channel:HTMLInputElement = document.getElementById("channel") as HTMLInputElement;
const button =document.querySelector("#submit")

const redirect=()=>{
    location.replace("http://"+window.location.host+"/channel/"+channel.value)
}
