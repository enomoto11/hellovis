import "./App.css";
import { useAuth0 } from "@auth0/auth0-react";

function App() {
  const { isAuthenticated, loginWithPopup, logout } = useAuth0();

  return (
    <div className="App">
      <header className="App-header">
        {!isAuthenticated ? (
          <button onClick={() => loginWithPopup()}>ログイン</button>
        ) : (
          <button
            onClick={() => {
              console.log("未実装");
              logout({});
            }}
          >
            Log out
          </button>
        )}
      </header>
    </div>
  );
}

export default App;
