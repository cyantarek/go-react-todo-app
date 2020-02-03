import React, {useState} from "react"
import {Icon} from "semantic-ui-react";
import TodoItem from "../TodoItem/TodoItem";

function TodoList(props) {
    let content = props.items.map((item, i) => {
        let color = "bg-light";
        let complete = "";

        if (item.status) {
            if (item.status === true) {
                color = "text-white bg-success";
                complete = "complete"
            } else {
                color = "text-white bg-danger";
            }
        }

        return <TodoItem {...props} key={i} item={item} color={color} complete={complete}/>
    });
    return (
        <div>
            {content}
        </div>
        /*<div>
            {this.state.items.map((item) => {
                let color = "bg-light";
                let complete = "";

                if (item.status) {
                    if (item.status === true) {
                        color = "text-white bg-success";
                        complete = "complete"
                    } else {
                        color = "text-white bg-danger";
                    }
                }
                return (

                );
            })}
        </div>*/
    );
}

export default TodoList