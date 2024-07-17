const UsernameErr = (props) => {
  return (
    <div
      className="alert alert-danger alert-dismissible fade show"
      role="alert"
    >
      Username Taken
      <button
        type="button"
        className="btn-close"
        data-bs-dismiss="alert"
        aria-label="Close"
        onClick={props.resetForm}
      ></button>
    </div>
  );
};

export default UsernameErr;
