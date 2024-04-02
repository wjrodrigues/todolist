import PropTypes from "prop-types";

Button.propTypes = {
  id: PropTypes.string,
  text: PropTypes.string,
  color: PropTypes.string,
  type: PropTypes.string,
};

function Button({ id, text, color, type }) {
  const btnClass = `text-white text-md font-bold rounded py-2 px-3 ${color}`;
  return (
    <>
      <button id={id} className={btnClass} type={type}>
        {text}
      </button>
    </>
  );
}

export default Button;
