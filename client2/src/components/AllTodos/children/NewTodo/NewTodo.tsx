const NewTodo: React.FC = () => {
	return (
		<div
			className="card m-3"
			style={{ width: "30%", backgroundColor: "aquamarine" }}
		>
			<div className="card-body">
				<h3 className="card-title">
					<input
						className="form-control"
						onChange={titleInputHandler}
						value={title}
						type="text"
					/>
				</h3>
				<hr />
				<div className="card-text">
					<textarea
						className="form-control mb-3"
						onChange={textInputHandler}
						value={text}
					/>
				</div>
				<hr />
				<div className="card-text">
					<button className="btn btn-info me-3 text-light" onClick={sendData}>
						Add
					</button>
					<button
						className="btn btn-danger text-light"
						onClick={props.onSubmit}
					>
						Cancel
					</button>
				</div>
			</div>
		</div>
	);
};

export default NewTodo;
