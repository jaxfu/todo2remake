import styles from "./App.module.scss";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Navbar from "../Navbar/Navbar";
import Footer from "../Footer/Footer";
import Register from "../Register/Register";
import Login from "../Login/Login";
import AllTodos from "../AllTodos/AllTodos";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { useState } from "react";
import { T_USER_DATA } from "../../types";

const queryClient = new QueryClient();

const App: React.FC = () => {
  const [userData, setUserData] = useState<T_USER_DATA>({ username: "", user_id: -1 });

  return (
    <QueryClientProvider client={queryClient}>
      <div className={styles.root}>
        <BrowserRouter>
          <Navbar username={userData.username} />
          <Routes>
            <Route path="/" element={<AllTodos userData={userData} />} />
            <Route path="/login" element={<Login setUserData={setUserData} />} />
            <Route path="/register" element={<Register setUserData={setUserData} />} />
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
