import styles from "./AddTodo.module.scss";

interface IProps {
	setAddingTodo: React.Dispatch<React.SetStateAction<boolean>>;
}

const AddTodo: React.FC<IProps> = (props) => {
	return (
		<div
			className={`card m-3 ${styles.root}`}
			onClick={() => {
				props.setAddingTodo(true);
			}}
		>
			<div className={`card-body`}>
				<h1 className="card-text">+</h1>
			</div>
		</div>
	);
};

export default AddTodo;
