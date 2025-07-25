# Lịch Âm Dương - Website

Website này sử dụng React + Vite + TypeScript để hiển thị thông tin lịch âm dương, chuyển đổi ngày, lịch tháng, ngày tốt/xấu.

## Tính năng
- Xem thông tin ngày hiện tại (dương/âm lịch, can chi, giờ hoàng đạo...)
- Chuyển đổi giữa dương lịch và âm lịch
- Xem lịch tháng (âm/dương)
- Tìm ngày tốt/xấu trong tháng
- Giao diện đẹp, hiện đại, responsive

## Cách chạy
1. Cài đặt dependencies:
   ```bash
   npm install
   ```
2. Chạy website:
   ```bash
   npm run dev
   ```
3. Truy cập: http://localhost:5173

## Kết nối API backend
- Website sẽ sử dụng các API backend đã có ở dự án để lấy dữ liệu lịch âm dương.
- Đảm bảo backend đã chạy ở http://localhost:8080

## Thư mục chính
- `src/`: Chứa mã nguồn React
- `.github/copilot-instructions.md`: Hướng dẫn cho Copilot

## Đóng góp
Mọi ý kiến đóng góp hoặc yêu cầu thêm tính năng vui lòng liên hệ qua Github.

# React + TypeScript + Vite

This template provides a minimal setup to get React working in Vite with HMR and some ESLint rules.

Currently, two official plugins are available:

- [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react) uses [Babel](https://babeljs.io/) for Fast Refresh
- [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react-swc) uses [SWC](https://swc.rs/) for Fast Refresh

## Expanding the ESLint configuration

If you are developing a production application, we recommend updating the configuration to enable type-aware lint rules:

```js
export default tseslint.config([
  globalIgnores(['dist']),
  {
    files: ['**/*.{ts,tsx}'],
    extends: [
      // Other configs...

      // Remove tseslint.configs.recommended and replace with this
      ...tseslint.configs.recommendedTypeChecked,
      // Alternatively, use this for stricter rules
      ...tseslint.configs.strictTypeChecked,
      // Optionally, add this for stylistic rules
      ...tseslint.configs.stylisticTypeChecked,

      // Other configs...
    ],
    languageOptions: {
      parserOptions: {
        project: ['./tsconfig.node.json', './tsconfig.app.json'],
        tsconfigRootDir: import.meta.dirname,
      },
      // other options...
    },
  },
])
```

You can also install [eslint-plugin-react-x](https://github.com/Rel1cx/eslint-react/tree/main/packages/plugins/eslint-plugin-react-x) and [eslint-plugin-react-dom](https://github.com/Rel1cx/eslint-react/tree/main/packages/plugins/eslint-plugin-react-dom) for React-specific lint rules:

```js
// eslint.config.js
import reactX from 'eslint-plugin-react-x'
import reactDom from 'eslint-plugin-react-dom'

export default tseslint.config([
  globalIgnores(['dist']),
  {
    files: ['**/*.{ts,tsx}'],
    extends: [
      // Other configs...
      // Enable lint rules for React
      reactX.configs['recommended-typescript'],
      // Enable lint rules for React DOM
      reactDom.configs.recommended,
    ],
    languageOptions: {
      parserOptions: {
        project: ['./tsconfig.node.json', './tsconfig.app.json'],
        tsconfigRootDir: import.meta.dirname,
      },
      // other options...
    },
  },
])
```
