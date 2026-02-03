"""
NeuralBlitz V50 - Intent Classification System
ML-based intent categorization using Random Forest classifier.
"""

try:
    from sklearn.ensemble import RandomForestClassifier
    from sklearn.model_selection import train_test_split, cross_val_score
    from sklearn.preprocessing import StandardScaler
    import joblib

    SKLEARN_AVAILABLE = True
except ImportError:
    SKLEARN_AVAILABLE = False

import numpy as np
from typing import List, Dict, Any, Optional, Tuple
from pathlib import Path
import json
from datetime import datetime

from ..minimal import MinimalCognitiveEngine, IntentVector


class IntentClassifier:
    """
    Machine Learning classifier for intent categorization.

    Automatically categorizes intents into cognitive styles:
    - creative: High innovation, moderate harmony
    - analytical: High knowledge, high dominance
    - social: High connection, high harmony
    - dominant: High control, low cooperation
    - balanced: All dimensions moderate
    - disruptive: High transformation, low preservation
    """

    CATEGORIES = [
        "creative",
        "analytical",
        "social",
        "dominant",
        "balanced",
        "disruptive",
    ]

    def __init__(self, model_path: Optional[str] = None):
        if not SKLEARN_AVAILABLE:
            raise ImportError(
                "scikit-learn not installed. "
                "Install with: pip install scikit-learn joblib"
            )

        self.model: Optional[RandomForestClassifier] = None
        self.scaler = StandardScaler()
        self.is_trained = False
        self.training_metadata: Dict[str, Any] = {}

        if model_path and Path(model_path).exists():
            self.load_model(model_path)

    def _extract_features(self, intent: IntentVector) -> np.ndarray:
        """Extract features from intent vector."""
        vector = intent.to_vector()

        # Basic features: the 7 dimensions
        features = vector.copy()

        # Derived features
        features = np.append(
            features,
            [
                np.mean(vector),  # Average intensity
                np.std(vector),  # Variability
                np.max(vector),  # Peak dimension
                np.min(vector),  # Minimum dimension
                vector[0] - vector[1],  # Dominance vs Harmony conflict
                vector[2] + vector[4],  # Change drive (creation + transformation)
                vector[3] + vector[6],  # Stability + Connection
            ],
        )

        return features

    def _auto_label(self, intent: IntentVector) -> str:
        """
        Auto-label intent based on heuristics.
        Used for initial training data generation.
        """
        vector = intent.to_vector()

        # Heuristic rules
        if vector[2] > 0.7 and vector[1] > 0.4:  # High creation, decent harmony
            return "creative"
        elif vector[5] > 0.7 and vector[0] > 0.5:  # High knowledge, decent dominance
            return "analytical"
        elif vector[6] > 0.7 and vector[1] > 0.6:  # High connection, high harmony
            return "social"
        elif vector[0] > 0.7 and vector[1] < 0.3:  # High dominance, low harmony
            return "dominant"
        elif (
            vector[4] > 0.7 and vector[3] < 0.3
        ):  # High transformation, low preservation
            return "disruptive"
        elif np.all(np.abs(vector) < 0.5):  # All moderate
            return "balanced"
        else:
            # Default to closest match
            scores = {
                "creative": vector[2] + vector[1] * 0.5,
                "analytical": vector[5] + vector[0] * 0.5,
                "social": vector[6] + vector[1],
                "dominant": vector[0] - vector[1],
                "disruptive": vector[4] - vector[3],
                "balanced": 1.0 - np.std(vector),
            }
            return max(scores, key=scores.get)

    def generate_training_data(
        self, n_samples: int = 1000, seed: int = 42
    ) -> Tuple[List[IntentVector], List[str]]:
        """
        Generate synthetic training data.

        Args:
            n_samples: Number of samples to generate
            seed: Random seed for reproducibility

        Returns:
            Tuple of (intents, labels)
        """
        np.random.seed(seed)

        intents = []
        labels = []

        for _ in range(n_samples):
            # Random intent with some structure
            intent = IntentVector(
                phi1_dominance=np.random.uniform(-1, 1),
                phi2_harmony=np.random.uniform(-1, 1),
                phi3_creation=np.random.uniform(-1, 1),
                phi4_preservation=np.random.uniform(-1, 1),
                phi5_transformation=np.random.uniform(-1, 1),
                phi6_knowledge=np.random.uniform(-1, 1),
                phi7_connection=np.random.uniform(-1, 1),
            )

            label = self._auto_label(intent)
            intents.append(intent)
            labels.append(label)

        return intents, labels

    def train(
        self,
        intents: List[IntentVector],
        labels: List[str],
        validation_split: float = 0.2,
        cross_validate: bool = True,
    ) -> Dict[str, Any]:
        """
        Train the classifier on labeled data.

        Args:
            intents: List of IntentVectors
            labels: List of category labels
            validation_split: Fraction for validation
            cross_validate: Whether to run cross-validation

        Returns:
            Training metrics
        """
        # Extract features
        X = np.array([self._extract_features(intent) for intent in intents])
        y = np.array(labels)

        # Scale features
        X_scaled = self.scaler.fit_transform(X)

        # Split data
        X_train, X_val, y_train, y_val = train_test_split(
            X_scaled, y, test_size=validation_split, random_state=42, stratify=y
        )

        # Train model
        self.model = RandomForestClassifier(
            n_estimators=100, max_depth=10, random_state=42, n_jobs=-1
        )

        self.model.fit(X_train, y_train)

        # Validate
        train_score = self.model.score(X_train, y_train)
        val_score = self.model.score(X_val, y_val)

        metrics = {
            "train_accuracy": train_score,
            "validation_accuracy": val_score,
            "n_samples": len(intents),
            "n_features": X.shape[1],
            "category_distribution": {cat: labels.count(cat) for cat in set(labels)},
        }

        # Cross-validation
        if cross_validate and len(intents) >= 100:
            cv_scores = cross_val_score(self.model, X_scaled, y, cv=5)
            metrics["cv_mean"] = cv_scores.mean()
            metrics["cv_std"] = cv_scores.std()

        self.is_trained = True
        self.training_metadata = {
            "timestamp": datetime.utcnow().isoformat(),
            "metrics": metrics,
            "n_samples": len(intents),
        }

        return metrics

    def predict(self, intent: IntentVector) -> Dict[str, Any]:
        """
        Predict category for an intent.

        Args:
            intent: IntentVector to classify

        Returns:
            Dictionary with prediction results
        """
        if not self.is_trained or self.model is None:
            return {"category": "unknown", "confidence": 0.0, "is_trained": False}

        # Extract features
        features = self._extract_features(intent)
        features_scaled = self.scaler.transform(features.reshape(1, -1))

        # Predict
        prediction = self.model.predict(features_scaled)[0]
        probabilities = self.model.predict_proba(features_scaled)[0]

        # Get confidence
        confidence = np.max(probabilities)

        # All category probabilities
        category_probs = dict(zip(self.model.classes_, probabilities))

        return {
            "category": prediction,
            "confidence": float(confidence),
            "all_probabilities": {k: float(v) for k, v in category_probs.items()},
            "is_trained": True,
        }

    def predict_batch(self, intents: List[IntentVector]) -> List[Dict[str, Any]]:
        """Predict categories for multiple intents."""
        return [self.predict(intent) for intent in intents]

    def get_feature_importance(self) -> Dict[str, float]:
        """Get feature importance from trained model."""
        if not self.is_trained or self.model is None:
            return {}

        feature_names = ["phi1", "phi2", "phi3", "phi4", "phi5", "phi6", "phi7"] + [
            "mean",
            "std",
            "max",
            "min",
            "dominance_harmony_conflict",
            "change_drive",
            "stability_connection",
        ]

        importances = self.model.feature_importances_

        return dict(zip(feature_names, importances.tolist()))

    def save_model(self, path: str) -> None:
        """Save trained model to disk."""
        if not self.is_trained:
            raise ValueError("Model not trained yet")

        model_data = {
            "model": self.model,
            "scaler": self.scaler,
            "categories": self.CATEGORIES,
            "metadata": self.training_metadata,
            "version": "50.0.0",
        }

        joblib.dump(model_data, path)

    def load_model(self, path: str) -> None:
        """Load trained model from disk."""
        model_data = joblib.load(path)

        self.model = model_data["model"]
        self.scaler = model_data["scaler"]
        self.training_metadata = model_data.get("metadata", {})
        self.is_trained = True

    def evaluate(
        self, test_intents: List[IntentVector], test_labels: List[str]
    ) -> Dict[str, Any]:
        """
        Evaluate model on test set.

        Returns detailed metrics including per-class performance.
        """
        from sklearn.metrics import classification_report, confusion_matrix

        predictions = [self.predict(intent)["category"] for intent in test_intents]

        report = classification_report(test_labels, predictions, output_dict=True)
        cm = confusion_matrix(test_labels, predictions, labels=self.CATEGORIES)

        return {
            "classification_report": report,
            "confusion_matrix": cm.tolist(),
            "accuracy": report["accuracy"],
            "macro_avg": report["macro avg"],
            "weighted_avg": report["weighted avg"],
        }


class IntentClassifierTrainer:
    """
    Helper class for training and managing the classifier.
    """

    def __init__(self, classifier: Optional[IntentClassifier] = None):
        self.classifier = classifier or IntentClassifier()
        self.engine = MinimalCognitiveEngine()

    def train_from_engine_patterns(
        self, n_samples: int = 1000, min_confidence: float = 0.7
    ) -> Dict[str, Any]:
        """
        Train classifier using patterns generated by the engine.

        This creates training data by processing diverse intents.
        """
        intents = []
        labels = []

        # Generate diverse intents
        for i in range(n_samples):
            # Create intent with varied characteristics
            intent = IntentVector(
                phi1_dominance=np.random.uniform(-1, 1),
                phi2_harmony=np.random.uniform(-1, 1),
                phi3_creation=np.random.uniform(-1, 1),
                phi4_preservation=np.random.uniform(-1, 1),
                phi5_transformation=np.random.uniform(-1, 1),
                phi6_knowledge=np.random.uniform(-1, 1),
                phi7_connection=np.random.uniform(-1, 1),
            )

            # Process through engine
            result = self.engine.process_intent(intent)

            # Only use high-confidence patterns
            if result["confidence"] >= min_confidence:
                # Auto-label
                label = self.classifier._auto_label(intent)
                intents.append(intent)
                labels.append(label)

        # Train
        metrics = self.classifier.train(intents, labels)

        return {
            "training_metrics": metrics,
            "n_samples_used": len(intents),
            "n_samples_requested": n_samples,
        }

    def create_default_model(self, save_path: str = "intent_classifier.joblib") -> None:
        """
        Create and save a default trained model.

        This can be used as a starting point for production.
        """
        # Generate training data
        intents, labels = self.classifier.generate_training_data(n_samples=2000)

        # Train
        metrics = self.classifier.train(intents, labels)

        print(f"Training complete:")
        print(f"  Train accuracy: {metrics['train_accuracy']:.2%}")
        print(f"  Validation accuracy: {metrics['validation_accuracy']:.2%}")
        if "cv_mean" in metrics:
            print(
                f"  CV accuracy: {metrics['cv_mean']:.2%} (+/- {metrics['cv_std']:.2%})"
            )

        # Save
        self.classifier.save_model(save_path)
        print(f"\n✅ Model saved to: {save_path}")


def quick_classify(intent: IntentVector, model_path: Optional[str] = None) -> str:
    """
    Quick classification function for one-off use.

    Args:
        intent: Intent to classify
        model_path: Optional path to saved model

    Returns:
        Category label
    """
    classifier = IntentClassifier(model_path)

    if not classifier.is_trained:
        # Train a quick model on-the-fly
        print("Training quick model...")
        trainer = IntentClassifierTrainer(classifier)
        trainer.create_default_model()

    result = classifier.predict(intent)
    return result["category"]


# Export
__all__ = ["IntentClassifier", "IntentClassifierTrainer", "quick_classify"]
