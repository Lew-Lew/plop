<template>
  <div class="home">
    <div class="d-flex justify-content-center">
      <div class="border-box" >
        <h1 id="title">Rettel</h1>
      </div>
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
        <li class="list-inline-item">{{this.userProp[9]}}</li>

      </ul>
    </div>

    <div id="get-letters-list-div">
      <ul class="list-inline" id="get-letters-list">
        <li v-on:click="getLetter(i)" v-for="(l, i) in letters"  v-bind:key="l" class="get-list-inline-item" >
          {{ l.letter.toUpperCase() }}
        </li>
      </ul>
    </div>

    <!-- <div v-show="this.result !== ''">
      <p v-if="this.result === true"> WIN </p>
      <p v-else> LOOSE </p>
    </div> -->

    <div v-show="this.result === ''" class="d-flex justify-content-center">
      <a v-on:click="validation()" type="button" class="btn btn-light buttons">Valider</a>
      <a v-on:click="deleteAll()" type="button" class="btn btn-light buttons">Effacer</a>
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
      result: '',
    }
  },

  methods: {
    getLetter(index){
      if (this.letters[index].checked === false) {
        this.userProp.push(this.letters[index].letter.toUpperCase())
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

  #title {
    font-family: 'Londrina Shadow', cursive;
    font-size: 75px;
    margin: 0px;
    color: #585757;
  }

  .border-box {
    border: 3px solid #585757;
    border-radius: 43px 43px 43px 0px;
    padding: 10px 60px;
  }

  #letters-list-div {
    margin: 30px;
  }


  #letters-list {
    display: flex;
    justify-content: center;
    align-items: center;
  }

  #letters-list li {
    width: 65px;
    height: 65px;
    padding: 20px;
    display: flex;
    justify-content: center;
    align-items: center;
    font-family: 'Londrina Solid', cursive;
    color: #585757;
    font-size: 35px;
    line-height: 40px;
    text-align: center;
    background: #fff;
    border: 2px solid #585757;
    border-radius: 9px 9px 9px 0px;
  }

  #get-letters-list-div {
    margin: 30px;
  }

  #get-letters-list {
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .get-list-inline-item {
    border-radius: 9px;
    background-color: #F8F8F8;
    font-family: 'Londrina Solid', cursive;
    font-size: 35px;
    color: #585757;
    width: 60px;
    height: 60px;
    padding: 20px;
    display: flex;
    justify-content: center;
    align-items: center;
    box-shadow: 0px 2px 4px -1px rgb(89 84 84 / 45%);
    margin-left: 5px;
    margin-right: 5px;
  }

  .get-list-inline-item:hover {
    border-radius: 9px;
    background-color: #F8F8F8;
    font-family: 'Londrina Solid', cursive;
    font-size: 35px;
    color: #585757;
    width: 60px;
    height: 60px;
    padding: 20px;
    display: flex;
    justify-content: center;
    align-items: center;
    box-shadow: 0px 2px 4px -1px rgb(89 84 84 / 45%) inset;
    margin-left: 5px;
    margin-right: 5px;
  }

  .active {
    background-color: red;
  }


  .buttons {
    font-family: 'Londrina Solid', cursive;
    font-size: 35px;
    background-color: #F8F8F8;
    box-shadow: 0px 2px 4px -1px rgb(89 84 84 / 45%);
    margin: 20px;
    width: 222px;
    height: 43px;
    display: flex;
    justify-content: center;
    align-items: center;
    border: none;
  }

  .buttons:hover {
    font-family: 'Londrina Solid', cursive;
    font-size: 35px;
    background-color: #F8F8F8;
    box-shadow: 0px 2px 4px -1px rgb(89 84 84 / 45%) inset;
    border: none;
    margin: 20px;
    width: 222px;
    height: 43px;
    display: flex;
    justify-content: center;
    align-items: center;
  }

</style>
