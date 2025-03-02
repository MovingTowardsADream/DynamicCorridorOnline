import style from "./ModalWindow.module.css"
import { useEffect } from "react";
import { FaTimes } from 'react-icons/fa';
import React from "react";

interface ModalProps {
    isOpen: boolean;
    onClose: () => void;
    message: string;
}

const Modal: React.FC<ModalProps> = ({ isOpen, onClose, message}) => {
    useEffect(() => {
        if (isOpen) {
            const timer = setTimeout(() => {
                onClose();
            }, 3000); // Закрыть через 3 секунды
            return () => clearTimeout(timer);
        }
    }, [isOpen, onClose]);

    if (!isOpen) return null;

    return (
        <div className={style.modalOverlay}>
            <div className={style.modalContent}>
                <button className={style.closeButton} onClick={onClose}>
                    <FaTimes color="#EF8354" size={16} /> {/* Иконка крестика */}
                </button>
                <p className={style.errorText}>{message}</p>
            </div>
        </div>
    );
};

export default Modal;