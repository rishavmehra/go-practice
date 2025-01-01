import React, { Component } from "react";
import './App.css';
import { connect, sentMsg } from './api';
import Header from './components/Header/Header';
import ChatHistory from './components/ChatHistory/ChatHistory';
import ChatInput from './components/ChatInput/ChatInput';

class App extends Component {

  constructor(props){
    super(props)
    this.state = {
      ChatHistory: []
    }
  }

  componentDidMount(){
    connect((msg)=>{
      console.log("New Message");
      this.setState(prevState=>({
        ChatHistory: [...this.state.ChatHistory, msg]
      }))
      console.log(this.state);
    })
  }

  send(event){
    if (event.keyCode===13)
    {
      sentMsg(event.target.value);
      event.target.value=""
    }
  }

  render(){
    return (
      <div className="App">
        <Header />
        <ChatHistory ChatHistory={this.state.ChatHistory}/>
        <ChatInput send={this.send}/>
        <button onClick={this.send}>Hit</button>
      </div>
    );
  }
}


export default App;
