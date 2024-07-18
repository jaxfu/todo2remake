import styles from "./App.module.scss";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Navbar from "../Navbar/Navbar";
import Footer from "../Footer/Footer";
import Register from "../Register/Register";
import Login from "../Login/Login";
import AllTodos from "../AllTodos/AllTodos";

const App: React.FC = () => {
	return (
		<div className={styles.root}>
			<BrowserRouter>
				<Navbar />
				<Routes>
					<Route path="/" element={<AllTodos />} />
					<Route path="/login" element={<Login />} />
					<Route path="/register" element={<Register />} />
					{/* <Route path="/about" element={<AboutMe />} />
					<Route path="/contact" element={<ContactInfo />} />
					<Route path="*" element={<ErrorPage />} /> */}
				</Routes>
				<Footer />
			</BrowserRouter>
			<div className={styles.push}></div>
		</div>
	);
};

export default App;
