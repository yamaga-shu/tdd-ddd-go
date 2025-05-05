# ユビキタス言語

## 貸出: Loan

| 日本語           | 英単語（システム用語） | 備考・メモ                                         |
| ---------------- | ---------------------- | -------------------------------------------------- |
| 貸出             | Loan                   | ある蔵書と会員の間の貸出行為                       |
| **フィールド**   |
| ID               | Loan.id                |                                                    |
| 借りた会員のID   | Loan.memberId          |                                                    |
| 借りられた蔵書ID | Loan.bookInventoryId   |                                                    |
| 貸出日           | Loan.loanDate          |                                                    |
| 返却予定日       | Loan.dueDate           | 貸出日 + 14日                                      |
| 返却日           | Loan.returnDate        | 実際に返却された日                                 |
| 延長の有無       | Loan.extended          | 延長は1度のみ                                      |
| 貸出状態         | Loan.status            | 次の貸出状態欄のいずれか                           |
| **貸出状態**     | LoanStatus             |
| 貸出中           | onLoan                 |                                                    |
| 返却済み         | returned               |                                                    |
| 延滞             | overdue                |                                                    |
| **操作**         |
| 登録             | Loan.Register          | 会員による蔵書の貸出の申請を受け、登録する操作     |
| 返却             | Loan.Return            | 会員による蔵書の貸出完了の申請を受け、返却する操作 |
| 延長申請         | Loan.Extend            | 会員による貸出期間延長の申請を受け、延長する操作   |
| **その他**       |
| 貸出冊数上限     | LoanLimit              | 5冊                                                |


## 会員: Member

| 日本語 | 英単語（システム用語） | 備考・メモ                         |
| ------ | ---------------------- | ---------------------------------- |
| 会員   | Member                 |                                    |
| 会員ID | MemberId               | システム内で会員を一意に識別       |
| 書籍   | Book                   | ISBNベースの一般的な書籍情報       |
| 蔵書   | BookInventory          | 個体ごとの本（同じISBNでも別管理） |
| 蔵書ID | BookInventoryId        | 個体識別用ID                       |

| 貸出履歴     | LoanHistory            | 過去の貸出記録                     |

| 延滞履歴     | OverdueHistory         | 延滞発生記録                       |

| 紛失報告     | LostReport             | 紛失したときの報告                 |
| 破損報告     | DamagedReport          | 破損したときの報告                 |
