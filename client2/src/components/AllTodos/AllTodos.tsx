import { useEffect, useState } from "react";

import Listings from "../components/Listings";
import AddTodo from "../components/AddTodo";
import NewTodo from "../components/NewTodo";
import type { T_TODO } from "../../types";
import { apiRequestGetTodos } from "../../methods/requests";
import type { AxiosResponse } from "axios";

const AllTodos: React.FC = () => {
	useEffect(() => {
		const res: AxiosResponse<T_TODO[]> = apiRequestGetTodos("");
		console.log(res.data);
	}, []);

	return <h1>AllTodos</h1>;
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

// 	return (
// 		<div>
// 			<div className="allTodosCont mt-3">
// 				{addingTodo ? (
// 					<NewTodo onSubmit={afterNewSubmit} />
// 				) : (
// 					<AddTodo onClick={addTodoMode} />
// 				)}
// 				{todos && <Listings data={todos} afterDelete={afterDelete} />}
// 			</div>
// 		</div>
// 	);
// };

export default AllTodos;
