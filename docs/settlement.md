# Settlement Algorithm

This project uses a greedy settlement algorithm.

## Approach

1. Calculate net balance of each user.
2. If balance > 0 → user will receive money.
3. If balance < 0 → user must pay money.
4. Match debtors with creditors.
5. Transfer minimum possible amount.

## Optimization

This minimizes the number of transactions.

## Complexity

Time Complexity: O(n)  
Space Complexity: O(n)