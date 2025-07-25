import React, { useEffect, useState } from 'react';

interface TodayInfo {
  solar: string;
  lunar: string;
  canChi: string;
  goodHours: string[];
  badHours: string[];
}

const API_BASE = 'http://localhost:8080/api/calendar';

const TodayCard: React.FC = () => {
  const [info, setInfo] = useState<TodayInfo | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetch(`${API_BASE}/today`)
      .then(res => {
        if (!res.ok) throw new Error('Không thể kết nối API');
        return res.json();
      })
      .then(data => {
        setInfo({
          solar: data.solarDate,
          lunar: data.lunarDate,
          canChi: data.canChi,
          goodHours: data.goodHours || [],
          badHours: data.badHours || [],
        });
        setLoading(false);
      })
      .catch(err => {
        setError(err.message || 'Lỗi không xác định');
        setLoading(false);
      });
  }, []);

  if (loading) return <div>Đang tải thông tin ngày...</div>;
  if (error) return <div style={{color: 'red'}}>Lỗi: {error}</div>;
  if (!info) return <div>Không lấy được dữ liệu ngày hôm nay.</div>;

  return (
    <div className="today-card">
      <h2>Thông tin ngày hôm nay</h2>
      <p><b>Dương lịch:</b> {info.solar}</p>
      <p><b>Âm lịch:</b> {info.lunar}</p>
      <p><b>Can Chi:</b> {info.canChi}</p>
      <p><b>Giờ hoàng đạo:</b> {info.goodHours.join(', ')}</p>
      <p><b>Giờ hắc đạo:</b> {info.badHours.join(', ')}</p>
    </div>
  );
};

export default TodayCard;
