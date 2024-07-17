import { Fragment } from "react";

import Todo from "./Todo";

const Listings = (props) => {
  return (
    <Fragment>
      {props.data.map((todo) => {
        return (
          <Todo
            key={todo._id}
            id={todo._id}
            title={todo.title}
            text={todo.text}
            afterDelete={props.afterDelete}
          />
        );
      })}
    </Fragment>
  );
};

export default Listings;
