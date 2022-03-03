import React,{Component} from 'react'
import axios from 'axios'
import {Card, Header, Form, Input, Icon} from 'semantic-ui-react'

let endpoint = "http://localhost:8080";

class ToDoList extends Component {
    constructor(props) {
        super(props);
        this.state = {
            newTodo: "",
            todos: []
        };
    }    
    ComponentDidMount() {
    this.getTask(); 
    }   
}

export default ToDoList;
