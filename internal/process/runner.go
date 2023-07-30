package process

import (
	"sync"

	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/tg"
)

type Runner struct {
	my                   *sync.Mutex
	processRepository    repository.ProcessRepository
	profileRepository    repository.ProfileRepository
	withdrawerRepository repository.WithdrawerRepository
	tgBot                *tg.Bot
	userRepository       repository.UserRepository
}

func NewRunner(
	processRepository repository.ProcessRepository,
	profileRepository repository.ProfileRepository,
	withdrawerRepository repository.WithdrawerRepository,
	tgBot *tg.Bot,
	userRepository repository.UserRepository,
) *Runner {
	return &Runner{
		my:                   &sync.Mutex{},
		processRepository:    processRepository,
		profileRepository:    profileRepository,
		withdrawerRepository: withdrawerRepository,
		tgBot:                tgBot,
		userRepository:       userRepository,
	}
}
