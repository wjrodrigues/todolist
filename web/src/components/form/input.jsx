import PropTypes from "prop-types";

Input.propTypes = {
  id: PropTypes.string,
  type: PropTypes.string,
  value: PropTypes.string,
  name: PropTypes.string,
  label: PropTypes.string,
  placeholder: PropTypes.string,
  required: PropTypes.bool
};

function Input({ id, type, value, name, label, placeholder, required }) {
  return (
    <>
      <label className="block text-gray-600 font-bold mb-2" htmlFor="name">
        {label}
      </label>

      <input
        className="appearance-none border border-gray-600 rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        id={id}
        type={type}
        name={name}
        value={value}
        placeholder={placeholder}
        required={required}
      />
    </>
  );
}

export default Input;
