import { useState } from "react";
import styles from "./NewTodo.module.scss";
import { T_TODO } from "../../../../types";

const NewTodo: React.FC = () => {
	const [todo, setTodo] = useState<T_TODO>({ id: 0, title: "", content: "" });

	function inputHandler(
		e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
	): void {
		setTodo((todo: T_TODO) => {
			return {
				...todo,
				[e.target.name]: e.target.value,
			};
		});
	}

	return (
		<div className={`card m-3 ${styles.root}`}>
			<div className="card-body">
				<h3 className="card-title">
					<input
						className="form-control"
						onChange={inputHandler}
						value={todo.title}
						name="title"
						type="text"
					/>
				</h3>
				<hr />
				<div className="card-text">
					<textarea
						className="form-control mb-3"
						onChange={inputHandler}
						name="content"
						value={todo.content}
					/>
				</div>
				<hr />
				<div className="card-text">
					<button className="btn btn-info me-3 text-light">Add</button>
					<button className="btn btn-danger text-light">Cancel</button>
				</div>
			</div>
		</div>
	);
};

export default NewTodo;
