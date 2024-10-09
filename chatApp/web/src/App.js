import logo  from './logo.svg';
import React, { Component } from "react";
import './App.css';
import { connect, sentMsg } from './api';

class App extends Component {

  constructor(props){
    super(props)
    connect()
  }

  send(){
    console.log("hello");
    sentMsg("Rishav loves to Code")
  }

  render(){
    return (
      <div className="App">
        <button onClick={this.send}>Hit</button>
      </div>
    );
  }
}


export default App;
