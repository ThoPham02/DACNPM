import "./input.css"
import { AiFillEyeInvisible } from "react-icons/ai"

const Input = ({ content, type, place, hasForgot, hasEyeInvis }) => {
    return (
        <div className="input">
            <div>
                <p>{content}</p>
                {hasForgot ? <p className="forgot-password">Quên mật khẩu</p> : <></>}
            </div>
            <input type={type} placeholder={place} />
            {hasEyeInvis ? <AiFillEyeInvisible /> : <></>}
        </div>
    )
}

export default Input;