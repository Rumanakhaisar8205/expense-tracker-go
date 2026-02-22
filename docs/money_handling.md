# Money Handling Strategy

All monetary values are stored using:

DECIMAL(10,2)

This avoids floating-point precision errors.

Example:
100.50
250.00
9999.99

PostgreSQL ensures accurate financial calculations.