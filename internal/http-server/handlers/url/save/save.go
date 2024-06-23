package save

import (
    "errors"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/render"
    "github.com/go-playground/validator/v10"
    "log/slog"
    "net/http"
    "url-sh/internal/lib/api/responce"
)

type Request struct {
    URL   string `json:"url" validate:"required,url"`
    Alias string `json:"alias,omitempty"`
}

type Responce struct {
    responce.Responce
    Alias string `json:"alias,omitempty"`
}

type URLSaver interface {
    SaveURL(urlToSave string, alias string) (int64, error)
}

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        const op = "handlers.url.save.New"

        log = log.With(
            slog.String("op", op),
            slog.String("request_id", middleware.GetReqID(r.Context())),
        )

        var req Request

        err := render.DecodeJSON(r.Body, &req)
        if err != nil {
            log.Error("failed to decode request body", err)

            render.JSON(w, r, responce.Error("failed to decode request"))

            return
        }

        log.Info("request body decoded", slog.Any("request", req))

        if err := validator.New().Struct(req); err != nil {
            var validateErr validator.ValidationErrors
            errors.As(err, &validateErr)

            log.Error("invalid request", err)

            render.JSON(w, r, responce.ValidationError(validateErr))

            return
        }

        // TODO: Alias logics

    }
}
