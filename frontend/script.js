const alphabet = ["a", "z", "e", "r", "t", "y", "u", "i", "o", "p", "m", "l", "k", "j", "h", "g", "f", "d", "s", "q", "w", "x", "c", "v", "b", "n"]
const letters_list = document.getElementById("letters-list")
const keyboard_list = document.getElementById("keyboard-list")
const table = document.querySelector("#table")

let array_letters = []
// si le mot fait 9 lettres, on prend aléatoirement 9 lettres:
for (let index = 0; index < 9; index++) {
  array_letters.push(alphabet[Math.floor(Math.random()*alphabet.length)])
}

// on affiche les lettres piochées
array_letters.forEach(l => {
  letters_list.insertAdjacentHTML("beforeend", `<li>${l}</li>`);
})

// style
for (const child of letters_list.children) {
  child.classList.add("list-inline-item");

  child.addEventListener("click", (event) => {
    letter = event.currentTarget.innerText;
    console.log(letter)
  })
}
