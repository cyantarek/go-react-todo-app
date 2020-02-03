import React, {useEffect, useState} from "react"
import 'bootstrap/dist/css/bootstrap.min.css';
import Header from "../../components/Header/Header.jsx";
import UserInput from "../../components/UserInput/UserInput";
import TodoList from "../../components/TodoList/TodoList";
import {MoonLoader} from "react-spinners";
import NoTodo from "../../components/NoTodo/NoTodo";

let api = "https://go-todo-backend.herokuapp.com";

function App(props) {
    const [userInput, setUserInput] = useState("");
    const [todoList, addItemToList] = useState([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState("");

    const handleUserInput = (e) => {
        setUserInput(e.target.value)
    };

    useEffect(() => {
        getAll();
    }, []);

    const getAll = () => {
        setLoading(true);
        fetch(api + "/api/todo").then(res => {
            if (res.status !== 200) {
                throw res.status
            } else {
                return res.json()
            }
        }).then(data => {
            if (data) {
                addItemToList([...data]);
            } else {
                addItemToList([])
            }
            setLoading(false)
        });
    };

    const handleTodoSubmit = (e) => {
        e.preventDefault();

        setLoading(true);

        if (userInput !== "") {
            fetch(api + "/api/todo", {
                method: "POST",
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({task: userInput})
            }).then(res => {
                if (res.status !== 200) {
                    throw res.status
                } else {
                    return res.json()
                }
            }).then(data => {
                addItemToList([...todoList, data]);
                setLoading(false);
                setUserInput("")
            })
        }
    };

    const updateTodo = (id) => {
        fetch(api + "/api/todo/" + id + "/complete", {
            method: "PUT",
            headers: {
                'Content-Type': 'application/json'
            }
        }).then(res => {
            if (res.status !== 200) {
                throw res.status
            } else {
                getAll()
            }
        });
    };

    const undoTodo = (id) => {
        fetch(api + "/api/todo/" + id + "/undo", {
            method: "PUT",
            headers: {
                'Content-Type': 'application/json'
            }
        }).then(res => {
            if (res.status !== 200) {
                throw res.status
            } else {
                getAll()
            }
        });
    };

    const deleteTodo = (id) => {
        fetch(api + "/api/todo/" + id, {
            method: "DELETE",
            headers: {
                'Content-Type': 'application/json'
            }
        }).then(res => {
            getAll()
        });
    };

    return (
        <div className="container pt-5">
            <div className="row">
                <div className="col-8 offset-2 text-center header">
                    <Header/>
                </div>
            </div>
            <div className="row mt-5">
                <div className={"container"}>
                    <div className="col-8 offset-2">
                        <UserInput value={userInput} onChange={handleUserInput} onSubmit={handleTodoSubmit}/>
                    </div>
                </div>
            </div>
            <div className="row">
                {loading ? (<div className="col-8 offset-2 d-flex justify-content-center">
                    <MoonLoader/>
                </div>) : (<div className="col-8 offset-2 d-block justify-content-center">
                    <TodoList items={todoList} handleUpdate={updateTodo} handleUndo={undoTodo}
                              handleDelete={deleteTodo}/>
                </div>)}
            </div>
            {todoList.length === 0 && !loading ? (<NoTodo/>) : null}
        </div>
    );
}

export default App