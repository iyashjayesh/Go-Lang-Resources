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
    
    componentDidMount() {
        this.getTask(); 
    }   

    render() {
        return (
            <div>
                <div className='row'>
                    <Header as='h2' icon textAlign='center' className='header'>
                        To Do List
                    </Header>
                </div>

                <div className='row'>
                    <Form onSubmit={this.addTask}>
                        <Input
                            type='text'
                            name='newTodo'
                            placeholder='Add a new task'
                            value={this.state.newTodo}
                            onChange={this.handleChange}
                            fluid
                        />
                        </Form>
                </div>

            </div>
        );
    }
}

export default ToDoList;
