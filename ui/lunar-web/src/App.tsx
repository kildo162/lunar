import './App.css';
import TodayCard from './components/TodayCard';

function App() {
  return (
    <div className="container">
      <h1>Lịch Âm Dương</h1>
      <TodayCard />
      {/* Các component khác sẽ được bổ sung: chuyển đổi lịch, lịch tháng, ngày tốt/xấu */}
    </div>
  );
}

export default App
