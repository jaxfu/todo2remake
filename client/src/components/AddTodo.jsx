import { Fragment } from "react";

const AddTodo = (props) => {
  return (
    <Fragment>
      <div
        className="card m-3"
        onClick={props.onClick}
        style={{ width: "30%", backgroundColor: "aquamarine" }}
      >
        <div className="card-body addTodo">
          <h1 className="card-text">+</h1>
        </div>
      </div>
    </Fragment>
  );
};

export default AddTodo;
