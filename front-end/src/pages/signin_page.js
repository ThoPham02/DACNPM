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
                                <p>Tên đăng nhập</p>
                            </div>
                            <input type="text" />
                        </div>
                        <div className="signin-input">
                            <div>
                                <p>Mật khẩu </p>
                                <p className="forgot-password">Quên mật khẩu</p>
                            </div>
                            <input type="text" />
                            <AiFillEyeInvisible />
                        </div>

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
                        Bạn chưa có tài khoản? <p>Đăng ký ngay!</p>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default SignIn;