const id = Math.random()
let sourse = new EventSource(`http://localhost:8080/notify?id=${id}`)


sourse.addEventListener('open', () => {
    console.log("OPEN:", id)
})

sourse.addEventListener('saludar', (event) => {
    console.log("SALUDAR:", event.data)
})

sourse.addEventListener('saltar', (event) => {
    console.log("SALTAR:", event.data)
})