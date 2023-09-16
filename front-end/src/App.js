import { useState } from 'react';
import './App.css';
import SignIn from './pages/signin_page';
import SignUp from './pages/signup_page';

function App() {
  const [page, setPage] = useState(true)

  return (
    <div className="App">
      {page ? <SignIn setPage={setPage} /> : <SignUp setPage={setPage} />}
    </div>
  );
}

export default App;
