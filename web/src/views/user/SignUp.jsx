import Input from "@/components/form/input.jsx";
import Button from "@/components/form/button.jsx";

function SignUp() {
  return (
    <div className="relative w-5/6 sm:w-2/4 xl:w-1/5 mx-auto top-36">
      <form className="w-26">
        <div className="mb-4">
          <Input id={"name"} type={"text"} label={"Nome"} required={true} />
        </div>

        <div className="mb-4">
          <Input id={"email"} type={"email"} label={"Email"} required={true} />
        </div>

        <div className="mb-4">
          <Input
            id={"password"}
            type={"password"}
            label={"Senha"}
            required={true}
          />
        </div>

        <div className="mb-4">
          <Button
            id={"create-account"}
            text={"Criar conta"}
            color={"bg-gray-950"}
            type={"submit"}
          />
        </div>
      </form>
    </div>
  );
}

export default SignUp;
