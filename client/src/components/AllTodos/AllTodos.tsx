import styles from "./AllTodos.module.scss";
import { apiRequestGetTodos } from "../../methods/requests";
import { useQuery } from "@tanstack/react-query";
import type { T_TODO, T_USER_DATA } from "../../types";
import Todo from "./children/Todo/Todo";
import AddTodo from "./children/AddTodo/AddTodo";
import { useState, useEffect } from "react";
import NewTodo from "./children/NewTodo/NewTodo";
import { useNavigate } from "react-router-dom";

interface IProps {
  userData: T_USER_DATA;
}

const AllTodos: React.FC<IProps> = (props) => {
  const navigate = useNavigate();
  useEffect(() => {
    if (props.userData.user_id === -1) {
      navigate("/login");
    }
  }, [])


  const [addingTodo, setAddingTodo] = useState<boolean>(false);

  const { isPending, isError, isSuccess, data } = useQuery({
    queryKey: ["todos"],
    queryFn: () => apiRequestGetTodos(props.userData.user_id),
  });

  if (isError) {
    console.log("Error");
    return <h2>Error...</h2>;
  }
  if (isPending) {
    return <h2>Loading...</h2>;
  }

  if (isSuccess) {
    console.log(`${JSON.stringify(data.data)}`);

    const todos: T_TODO[] = data.data;
    return (
      <div className={`${styles.root}`}>
        {addingTodo ? (
          <NewTodo setAddingTodo={setAddingTodo} />
        ) : (
          <AddTodo setAddingTodo={setAddingTodo} />
        )}

        {todos.length > 0 && todos.map((todo) => (
          <Todo key={todo.id} todo={todo} />
        ))}
      </div>
    );
  }


};

// const AllTodos = () => {
// 	const [todos, setTodos] = useState<T_TODO[]>([]);
// 	const [addingTodo, setAddingTodo] = useState(false);
// 	const [rerender, setRerender] = useState(false);

// 	useEffect(() => {
// 		fetchData();
// 		console.log("UseEffect Triggered");
// 	}, [addingTodo, rerender]);

// 	const fetchData = () => {
// 		try {
// 			fetch("http://localhost:5000/api")
// 				.then((res) => {
// 					return res.json();
// 				})
// 				.then((data) => {
// 					setTodos(data);
// 				});
// 		} catch (e) {
// 			console.log(`Error: ${e}`);
// 		}
// 	};

// 	const addTodoMode = () => {
// 		if (addingTodo === false) {
// 			setAddingTodo(true);
// 		}

// 		return;
// 	};

// 	const afterNewSubmit = () => {
// 		setAddingTodo(false);
// 	};

// 	const afterDelete = () => {
// 		rerender ? setRerender(false) : setRerender(true);
// 	};

// return (
//   <div>
//     <div className="allTodosCont mt-3">
//       {addingTodo ? (
//         <NewTodo onSubmit={afterNewSubmit} />
//       ) : (
//         <AddTodo onClick={addTodoMode} />
//       )}
//       {todos && <Listings data={todos} afterDelete={afterDelete} />}
//     </div>
//   </div>
// );
//  };

export default AllTodos;
