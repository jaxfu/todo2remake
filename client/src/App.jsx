import { BrowserRouter, Routes, Route } from "react-router-dom";

import Navbar from "./layouts/Navbar";
import Footer from "./layouts/Footer";
import AllTodos from "./views/AllTodos";
import AboutMe from "./views/AboutMe";
import ContactInfo from "./views/ContactInfo";
import LoginPage from "./views/LoginPage";
import ErrorPage from "./views/ErrorPage";
import RegisterPage from "./views/RegisterPage";

const App = () => {
  return (
    <div className="wrapper">
      <BrowserRouter>
        <Navbar />
        <Routes>
          <Route path="/" element={<AllTodos />} />
          <Route path="/about" element={<AboutMe />} />
          <Route path="/contact" element={<ContactInfo />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/register" element={<RegisterPage />} />
          <Route path="*" element={<ErrorPage />} />
        </Routes>
        <Footer />
      </BrowserRouter>
      <div className="push"></div>
    </div>
  );
};

export default App;
