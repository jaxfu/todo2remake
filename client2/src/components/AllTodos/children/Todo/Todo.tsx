import { useState } from "react";
import type { T_TODO } from "../../../../types";

interface IProps {
	todo: T_TODO;
}

const Todo: React.FC<IProps> = (props) => {
	const [editMode, setEditMode] = useState(false);
	const [title, setTitle] = useState(props.todo.title);
	const [content, setContent] = useState(props.todo.content);

	const titleChangeHandler = (e: any) => {
		setTitle(e.target.value);
	};

	const textChangeHandler = (e: any) => {
		setContent(e.target.value);
	};

	// const sendDelete = async () => {
	// 	const { id } = props;

	// 	if (!editMode) {
	// 		try {
	// 			await fetch(`http://localhost:5000/api/${id}`, {
	// 				method: "DELETE",
	// 			}).then((res) => res.json());
	// 		} catch (e) {
	// 			console.log(`Error: ${e}`);
	// 		}
	// 	} else if (editMode) {
	// 		setEditMode(false);
	// 	}
	// };

	// const sendPut = async (data) => {
	// 	try {
	// 		await fetch("http://localhost:5000/api", {
	// 			method: "PUT",
	// 			headers: {
	// 				"Content-Type": "application/json",
	// 			},
	// 			body: JSON.stringify(data),
	// 		});
	// 		return;
	// 	} catch (e) {
	// 		console.log(`Error: ${e}`);
	// 	}
	// };

	const toggleEdit = () => {
		if (editMode === true) {
			const sendData = {
				id: props.todo.id,
				title,
				text: content,
			};

			//sendPut(sendData);
		}

		setEditMode(editMode ? false : true);
	};

	return (
		<div className="card m-3 todo">
			<div className="card-body">
				<h3 className="card-title">
					{editMode ? (
						<input
							className="form-control"
							onChange={titleChangeHandler}
							type="text"
							value={title}
						/>
					) : (
						title
					)}
				</h3>
				<hr />
				<div className="card-text">
					{editMode ? (
						<textarea
							className="form-control"
							onChange={textChangeHandler}
							value={content}
						></textarea>
					) : (
						content
					)}
				</div>
				<hr />
				<div className="card-text">
					<button className="btn btn-info me-3 text-light">
						{editMode ? "Submit" : "Edit"}
					</button>
					<button className="btn btn-secondary text-light">
						{editMode ? "Cancel" : "Delete"}
					</button>
				</div>
			</div>
		</div>
	);
};

export default Todo;
