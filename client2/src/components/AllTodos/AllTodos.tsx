import { apiRequestGetTodos } from "../../methods/requests";
import { useQuery } from "@tanstack/react-query";

const AllTodos: React.FC = () => {
	const { isError, isSuccess, data } = useQuery({
		queryKey: ["todos"],
		queryFn: () => apiRequestGetTodos(""),
	});

	if (isError) console.log("Error");
	if (isSuccess) console.log(`${JSON.stringify(data.data)}`);

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
