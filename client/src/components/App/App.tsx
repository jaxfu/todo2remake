import styles from "./App.module.scss";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Navbar from "../Navbar/Navbar";
import Footer from "../Footer/Footer";
import Register from "../Register/Register";
import Login from "../Login/Login";
import AllTodos from "../AllTodos/AllTodos";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { useState } from "react";

const queryClient = new QueryClient();

const App: React.FC = () => {
  const [username, setUsername] = useState<string>("");

  return (
    <QueryClientProvider client={queryClient}>
      <div className={styles.root}>
        <BrowserRouter>
          <Navbar />
          <Routes>
            <Route path="/" element={<AllTodos username={username} />} />
            <Route path="/login" element={<Login setUsername={setUsername} />} />
            <Route path="/register" element={<Register setUsername={setUsername} />} />
            {/* <Route path="/about" element={<AboutMe />} />
					<Route path="/contact" element={<ContactInfo />} />
					<Route path="*" element={<ErrorPage />} /> */}
          </Routes>
          <Footer />
        </BrowserRouter>
        <div className={styles.push}></div>
      </div>
    </QueryClientProvider>
  );
};

export default App;
