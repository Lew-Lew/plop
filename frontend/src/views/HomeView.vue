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

    <div v-show="this.result !== ''">
      <p v-if="this.result === true"> WIN </p>
      <p v-else> LOOSE </p>
    </div>

    <div v-show="this.result === ''">
      <button v-on:click="deleteAll()" type="button" class="btn btn-light">Effacer</button>

      <button v-on:click="validation()" type="button" class="btn btn-light">Valider</button>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import axios from 'axios';

export default {
  name: 'HomeView',
  data() {
    return {
      letters: [],
      userProp: [],
      result: ''
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
      const headers = {
        'Content-Type':'text/plain'
      };
      axios
        .post("http://localhost:8081/word", {"word": this.userProp.join("")},{headers})
        .then(response => {
          this.result = response.data.exist
        })
    }
  },

  mounted () {
    axios
      .get('http://localhost:8081/word')
      .then(response => {
        for (const resp of response.data) {
          this.letters.push({letter: resp, checked:false})
        }
      })
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
