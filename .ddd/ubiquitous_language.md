# ユビキタス言語

| 日本語       | 英単語（システム用語） | 備考・メモ                         |
| ------------ | ---------------------- | ---------------------------------- |
| 会員         | Member                 |                                    |
| 会員ID       | MemberId               | システム内で会員を一意に識別       |
| 書籍         | Book                   | ISBNベースの一般的な書籍情報       |
| 蔵書         | BookInventory          | 個体ごとの本（同じISBNでも別管理） |
| 蔵書ID       | BookInventoryId        | 個体識別用ID                       |
| 貸出         | Loan                   | ある蔵書と会員の間の貸出行為       |
| 貸出登録     | LoanRegistration       | 貸出をリクエスト・登録する操作     |
| 貸出履歴     | LoanHistory            | 過去の貸出記録                     |
| 貸出冊数上限 | LoanLimit              | 5冊                                |
| 返却予定日   | DueDate                | 貸出日＋14日                       |
| 延滞         | Overdue                | 返却予定日を過ぎて未返却の状態     |
| 延滞履歴     | OverdueHistory         | 延滞発生記録                       |
| 延長申請     | LoanExtensionRequest   | 貸出期間延長の申請                 |
| 貸出停止     | LoanSuspension         | 延滞などによる貸出一時停止状態     |
| 紛失報告     | LostReport             | 紛失したときの報告                 |
| 破損報告     | DamagedReport          | 破損したときの報告                 |
