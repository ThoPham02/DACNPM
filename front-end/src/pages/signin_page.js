import "./login.css"
import { AiOutlineClose } from "react-icons/ai"
import { FcGoogle } from "react-icons/fc"
import { FaFacebookF } from "react-icons/fa"
import Input from "../components/input/input"

const SignIn = ({setPage}) => {
    const handleChangepage = () => {
        setPage(false)
    }

    return (
        <div className="container">
            <div className="signin-block">
                <div className="signin-header">
                    <h2>Đăng nhập</h2>
                    <AiOutlineClose />
                </div>
                <div className="signin-body">
                    <div className="signin-form">
                        <Input content="Tên đăng nhập" type="text" place="" />
                        <Input content="Mật khẩu" type="text" hasForgot={true} hasEyeInvis={true} />
                        <button className="signin-button">
                            <p>Đăng nhập</p>
                        </button>
                    </div>

                    <div className="signin-other">
                        <p>hoặc đăng nhập bằng</p>
                        <div className="signin-method">
                            <FaFacebookF /> <p>Facebook</p>
                        </div>
                        <div className="signin-method">
                            <FcGoogle /> <p>Google</p>
                        </div>
                    </div>

                    <div className="signin-out">
                        Bạn chưa có tài khoản? <p onClick={handleChangepage}>Đăng ký ngay!</p>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default SignIn;