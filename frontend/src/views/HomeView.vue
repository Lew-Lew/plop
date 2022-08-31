<template>
  <div class="home">
    <h1>PLOP</h1>
    <div id="letters-list-div">
      <ul class="list-inline" id="letters-list">
        <li v-on:click="getLetter(i)" v-for="(l, i) in letters"  v-bind:key="l" class="list-inline-item">
          {{ l.letter }}
        </li>
      </ul>
    </div>

    <div id="letters-list-div">
      <ul class="list-inline" id="letters-list">
        <li class="list-inline-item">{{this.userProp[0]}}</li>
        <li class="list-inline-item">{{this.userProp[1]}}</li>
        <li class="list-inline-item">{{this.userProp[2]}}</li>
        <li class="list-inline-item">{{this.userProp[3]}}</li>
        <li class="list-inline-item">{{this.userProp[4]}}</li>
        <li class="list-inline-item">{{this.userProp[5]}}</li>
        <li class="list-inline-item">{{this.userProp[6]}}</li>
        <li class="list-inline-item">{{this.userProp[7]}}</li>
        <li class="list-inline-item">{{this.userProp[8]}}</li>
      </ul>
    </div>

    <div>
      <button v-on:click="deleteAll()" type="button" class="btn btn-light">Effacer</button>

      <button v-on:click="validation()" type="button" class="btn btn-light">Valider</button>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src

export default {
  name: 'HomeView',
  data() {
    return {
      letters: [
        {letter: "b", checked: false},
        {letter: "a", checked: false},
        {letter: "n", checked: false},
        {letter: "a", checked: false},
        {letter: "n", checked: false},
        {letter: "i", checked: false},
        {letter: "e", checked: false},
        {letter: "r", checked: false},
        {letter: "s", checked: false}
      ],
      userProp: []
    }
  },

  methods: {
    getLetter(index){
      if (this.letters[index].checked === false) {
        this.userProp.push(this.letters[index].letter)
        this.letters[index].checked = true
      }
    },

    deleteAll(){
      this.userProp = []
      this.letters.forEach(element => {
        element.checked = false
      });
    },

    validation(){
      console.log(this.userProp.join(""))
    }
  }
}
</script>

<style scoped>
#letters-list {
  display: flex;
  justify-content: center;
  align-items: center;
}

#letters-list li {
  width: 40px;
  height: 40px;
  line-height: 40px;
  text-align: center;
  background: #fff;
  border: 1px solid #f9f9f9;
  border-radius: 5px;
}
</style>
