package registration

type RegistrationService struct {
	notifier RegistrationNotifier
}

func NewRegistrationService(notifier RegistrationNotifier) *RegistrationService {
	return &RegistrationService{
		notifier,
	}
}

func (rs *RegistrationService) HandleNewRegistration(registration *Registration) error {
	return rs.notifier.InformAboutNewRegistration(registration)
}
