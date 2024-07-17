const InvalidFormData = (props) => {
  return (
    <div
      className="alert alert-danger alert-dismissible fade show mt-3"
      role="alert"
    >
      Please Enter A Valid Username And Password
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

export default InvalidFormData;
