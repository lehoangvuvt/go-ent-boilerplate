package transactiondomain

import "time"

func (t *Transaction) Confirm() error {
	if t.Status != Pending {
		return ErrInvalidTransition
	}

	t.Status = Completed
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Transaction) Cancel() error {
	if t.Status != Pending {
		return ErrInvalidTransition
	}

	t.Status = Rejected
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Transaction) Fail() error {
	if t.Status != Pending {
		return ErrInvalidTransition
	}

	t.Status = Failed
	t.UpdatedAt = time.Now()
	return nil
}
