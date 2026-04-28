# 🧠 GoGym — Definição de Domínio (MVP)

## 🎯 Objetivo
Sistema para rastrear cargas de treino de musculação, permitindo ajustar progressivamente peso, séries e repetições dos exercícios ao longo do tempo.

---

## 🧱 Entidades

### 🏋️ Workout (Treino)
Representa um conjunto de exercícios.

#### Propriedades:
- `id`
- `name`

#### Regras de negócio:
- Deve possuir um nome
- Não pode conter exercícios duplicados

---

### 💪 Exercise (Exercício)
Representa um exercício padrão do sistema.

#### Propriedades:
- `id`
- `name`
- `muscle_group` (enum)

#### Regras de negócio:
- Nome deve ser único
- Deve possuir grupo muscular definido

---

### 🔗 WorkoutExercise
Representa um exercício dentro de um treino.

#### Propriedades:
- `id`
- `workout_id`
- `exercise_id`
- `sets`
- `reps`
- `weight`
- `last_weight_update_at`

#### Regras de negócio:
- `sets` >= 1
- `reps` >= 1
- `weight` > 0
- Um exercício não pode se repetir dentro do mesmo treino
- Peso não varia por série (regra atual do sistema)
- `last_weight_update_at` deve ser atualizado sempre que o peso for alterado

---

## 🔗 Relacionamentos

- Um `Workout` possui vários `WorkoutExercise`
- Um `Exercise` pode estar em vários `WorkoutExercise`

---

## 🧠 Decisões de Design

- O sistema não controla execução em tempo real (sem histórico por enquanto)
- O foco é apenas atualização manual de carga
- Não existe ordem fixa de exercícios dentro do treino
- Grupo muscular será representado como enum (evitando strings soltas)
- Peso não varia por série (simplificação intencional)

---