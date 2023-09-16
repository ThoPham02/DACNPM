import Input from "../components/input/input";
import "./login.css"

import { FaFacebookF } from "react-icons/fa"
import { AiFillApple } from "react-icons/ai"
import { BiLogoGooglePlus } from "react-icons/bi"

const SignUp = ({ setPage }) => {
    const handleChangePage = () => {
        setPage(true)
    }

    return (
        <div className="container">
            <div className="signup-block">
                <div className="signup-header">
                    <h2>Đăng ký tài khoản mới miến phí</h2>
                </div>
                <div className="signup-body">
                    <div className="signup-form">
                        <Input content="Họ và tên" type="text" place="Nhập họ và tên của bạn" />
                        <Input content="Tên đăng nhập" type="text" place="Nhập tên đăng nhập của bạn" />
                        <Input content="Mật khẩu" type="password" place="Nhập mật khẩu của bạn" hasEyeInvis={true} />
                        <Input content="Nhập lại mật khẩu" type="password" place="Nhập lại mật khẩu của bạn" hasEyeInvis={true} />
                        <Input content="Email" type="email" place="Nhập email của bạn" />
                        <div>
                            <input type="checkbox" /> Tôi đồng ý với các <span>điều kiện và điều khoản</span> của ThọPhạm
                        </div>
                        <button className="signup-button">
                            <p>Đăng ký</p>
                        </button>
                    </div>

                    <div className="signup-other">
                        <div >
                            <p>Hoặc</p>
                            <span className="line"></span>
                        </div>
                        <div className="method-container">
                            <div className="signup-method">
                                <FaFacebookF />
                            </div>
                            <div className="signup-method">
                                <BiLogoGooglePlus />
                            </div>
                            <div className="signup-method">
                                <AiFillApple />
                            </div>
                        </div>
                    </div>

                    <div className="signup-out">
                        Bạn đã có tài khoản rồi? {" "} <p onClick={handleChangePage}>Đăng nhập</p>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default SignUp;