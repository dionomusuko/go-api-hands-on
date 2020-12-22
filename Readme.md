# API ハンズオン

## Goの基本ライブラリのみを使ってTodoアプリを作る

### 使用ライブラリ
- FW : Echo
- ORM : database/sql
- LiveReload : air

### エンドポイント
```
GET /tasks  　タスク一覧取得
GET /tasks/:id　 idに該当するタスクを取得
POST /tasks 　タスク新規作成
PUT /tasks 　タスク内容変更
DELETE /tasks/:id 　タスク削除
```

### モデル
| Field | type | description |
| :--- | :--- | :--- | 
| id | int | auto incrementすること |
| title | varchar(256) | タスクのタイトル |
| description | varchar(256) | タスクの説明 |
| created_at | datetime ||
| updated_at | datetime ||
