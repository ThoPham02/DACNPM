import "./login.css"
import { AiOutlineClose, AiFillEyeInvisible } from "react-icons/ai"
import { FcGoogle } from "react-icons/fc"
import { FaFacebookF } from "react-icons/fa"

const SignIn = (setPage) => {
    return (
        <div className="container">
            <div className="signin-block">
                <div className="signin-header">
                    <h2>Đăng nhập</h2>
                    <AiOutlineClose />
                </div>
                <div className="signin-body">
                    <div className="signin-form">
                        <div className="signin-input">
                            <div>
                                <span>Tên đăng nhập</span>
                            </div>
                            <input type="text" />
                        </div>
                        <div className="signin-input">
                            <div>
                                <span>Mật khẩu </span>
                                <span>Quên mật khẩu</span>
                            </div>
                            <input type="text" />
                            <AiFillEyeInvisible />
                        </div>

                        <button className="signin-button">Đăng nhập</button>
                    </div>

                    <div className="signin-other">
                        <p>hoặc đăng nhập bằng</p>
                        <div className="signin-method">
                            <FaFacebookF /> Facebook
                        </div>
                        <div className="signin-method">
                            <FcGoogle /> Google
                        </div>
                    </div>

                    <div className="signin-out">
                        Bạn chưa có tài khoản? <span>Đăng ký ngay!</span>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default SignIn;