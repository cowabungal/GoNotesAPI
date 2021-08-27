package service

import (
	GoNotes "GoNotes"
	"GoNotes/pkg/repository"
)

type NoteService struct {
	repo *repository.Repository
}

func NewNoteService(repo *repository.Repository) *NoteService {
	return &NoteService{repo: repo}
}

func (s *NoteService) GetAll(userId int) ([]*GoNotes.Note, error) {
	notes, err := s.repo.Note.GetAll(userId)
	return notes, err
}

func (s *NoteService) Add(id int, note *GoNotes.Note) (int, error) {
	return s.repo.Note.Add(id, note)
}

func (s *NoteService) Delete(id, userId int) error {
	return s.repo.Delete(id, userId)
}
